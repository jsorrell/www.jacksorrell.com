'use strict';

var path = require('path');

var autoprefixer = require('autoprefixer');
var browserSync = require('browser-sync').create();
var browserify = require('browserify');
var spawn = require('child_process').spawn;
var gulp = require('gulp');
const gulpClean = require('gulp-clean');
var minifycss = require('gulp-clean-css');
const happiness = require('gulp-happiness');
const htmlmin = require('gulp-htmlmin');
var newer = require('gulp-newer');
var postcss = require('gulp-postcss');
var print = require('gulp-print').default;
var gulpSass = require('gulp-sass');
var sourcemaps = require('gulp-sourcemaps');
var tap = require('gulp-tap');
const gulpTslint = require('gulp-tslint').default;
var typescript = require('gulp-typescript');
var uglify = require('gulp-uglify-es').default;
const gulpWatchSass = require('gulp-watch-sass');
var tsify = require('tsify');
var tslint = require('tslint');
var buffer = require('vinyl-buffer');

const yaml = require('js-yaml');
const fs = require('fs');
const argv = require('yargs').argv;

var serverPort = argv.port;

if (serverPort === undefined) {
	var configFile = (argv.config === undefined) ? 'config.yaml' : argv.config;
	try {
		serverPort = yaml.safeLoad(fs.readFileSync(configFile, 'utf8')).server.port;
	} catch (e) {
		serverPort = 3000;
	}
}

var bsPort = (argv.bsport === undefined) ? serverPort + 1 : Number(argv.bsport);

const sass = { src: path.join(__dirname, 'client/sass/'), dest: path.join(__dirname, 'assets/public/css/') };
const goHtml = { src: path.join(__dirname, 'client/views/'), dest: path.join(__dirname, 'assets/templates/') };
const ts = { src: path.join(__dirname, 'client/scripts/'), dest: path.join(__dirname, 'assets/public/js/') };
const favicon = { src: path.join(__dirname, 'client/favicon/public/'), dest: path.join(__dirname, 'assets/public/fav/') };

var clientTs = typescript.createProject(path.join(ts.src, 'tsconfig.json'));

export function compileSass (stream) {
	if (!stream || !stream.pipe) stream = gulp.src(path.join(sass.src, '**/*.scss'));

	return stream
	.pipe(sourcemaps.init())
	.pipe(gulpSass().on('error', gulpSass.logError))
	.pipe(postcss([autoprefixer({
		cascade: false
	})]))
	.pipe(minifycss())
	.pipe(print(filepath => `built: ${filepath}`))
	.pipe(sourcemaps.write('.'))
	.pipe(gulp.dest(sass.dest));
}

export function compileTs () {
	return clientTs.src()
	.pipe(tap(function (file) {
		file.path = file.path.slice(0, -3) + '.js';
	}))
	.pipe(newer(ts.dest))
	.pipe(tap(function (file) {
		file.path = file.path.slice(0, -3) + '.ts';
		file.contents = browserify(file.path, {debug: true})
			.plugin(tsify, {project: path.join(ts.src, 'tsconfig.json')})
			.bundle();
		file.path = file.path.slice(0, -3) + '.js';
	}))
	.pipe(buffer())
	.pipe(sourcemaps.init({ loadMaps: true }))
	.pipe(uglify())
	.pipe(print(filepath => `built: ${filepath}`))
	.pipe(sourcemaps.write('.'))
	.pipe(gulp.dest(ts.dest));
}

// TODO generate favicons
export function copyFavicons () {
	return gulp.src(path.join(favicon.src, '**'))
		.pipe(newer(favicon.dest))
		.pipe(print(filepath => `copied: ${filepath}`))
		.pipe(gulp.dest(favicon.dest));
}

export function copyViews () {
	return gulp.src(path.join(goHtml.src, '**/*.gohtml'))
		.pipe(newer(goHtml.dest))
		.pipe(htmlmin({ collapseWhitespace: true }))
		.pipe(print(filepath => `copied: ${filepath}`))
		.pipe(gulp.dest(goHtml.dest));
}

export const build = gulp.parallel(compileSass, compileTs, copyFavicons, copyViews);

function bsInit (done) {
	browserSync.init({
		proxy: 'localhost:' + serverPort,
		port: bsPort,
		ui: {
			port: bsPort + 1
		},
		open: false
	});
	done();
}

function reloadBs (done) {
	browserSync.reload();
	done();
}

export function watch () {
	gulp.watch(path.join(ts.src, '**/*.ts'), gulp.series(compileTs, reloadBs));
	gulp.watch(path.join(goHtml.src, '**/*.gohtml'), gulp.series(copyViews, reloadBs));
	compileSass(gulpWatchSass(path.join(sass.src, '**/*.scss'))).pipe(browserSync.stream());
}

export const dev = gulp.series(build, bsInit, watch);

export function lintJs () {
	return gulp.src(['**/*.js', '!node_modules/**', '!' + path.join(ts.dest, '**')])
		.pipe(happiness())
		.pipe(happiness.format())
		.pipe(happiness.failAfterError());
}

export function lintTs () {
	const program = tslint.Linter.createProgram(path.join(ts.src, 'tsconfig.json'));
	return gulp.src([path.join(ts.src, '**/*.ts')])
		.pipe(gulpTslint({
			formatter: 'verbose',
			program: program,
			configuration: './tslint.json'
		}))
		.pipe(gulpTslint.report());
}

export function lintGo () {
	return spawn('goimports', ['-w', '.'], {
		stdio: 'inherit'
	});
}

export const lint = gulp.series(lintJs, lintTs, lintGo);

export function cleanTemplates () {
	return gulp.src(goHtml.dest, { read: false })
		.pipe(gulpClean());
}

export function cleanCss () {
	return gulp.src(sass.dest, { read: false })
		.pipe(gulpClean());
}

export function cleanJs () {
	return gulp.src(ts.dest, { read: false })
		.pipe(gulpClean());
}

export function cleanFavicons () {
	return gulp.src(favicon.dest, { read: false })
		.pipe(gulpClean());
}

export const clean = gulp.parallel(cleanTemplates, cleanCss, cleanJs, cleanFavicons);

export default build;
