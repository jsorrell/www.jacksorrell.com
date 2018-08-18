'use strict';

require('dotenv').config();
var gulp = require('gulp');
var autoprefixer = require('autoprefixer');
var browserSync = require('browser-sync').create();
const happiness = require('gulp-happiness');
var minifycss = require('gulp-clean-css');
var postcss = require('gulp-postcss');
var gulpSass = require('gulp-sass');
var spawn = require('child_process').spawn;
var typescript = require('gulp-typescript');
var tsify = require('tsify');
var browserify = require('browserify');
var uglify = require('gulp-uglify-es').default;
var buffer = require('vinyl-buffer');
var newer = require('gulp-newer');
var print = require('gulp-print').default;
const gulpTslint = require('gulp-tslint').default;
var tslint = require('tslint');
const gulpWatchSass = require('gulp-watch-sass');
var sourcemaps = require('gulp-sourcemaps');
var tap = require('gulp-tap');

var serverTs = typescript.createProject('src/tsconfig.json');
var clientTs = typescript.createProject('src/scripts/tsconfig.json');

function compileSass (stream) {
	return stream
	.pipe(sourcemaps.init())
	.pipe(gulpSass().on('error', gulpSass.logError))
	.pipe(postcss([autoprefixer({
		cascade: false
	})]))
	.pipe(minifycss())
	.pipe(print(filepath => `built: ${filepath}`))
	.pipe(sourcemaps.write('.'))
	.pipe(gulp.dest('dist/public/css/'));
}

export function sass () {
	return compileSass(gulp.src('src/sass/style.scss'));
}

function compileClientTs () {
	return clientTs.src()
	.pipe(tap(function (file) {
		file.contents = browserify(file.path, {debug: true})
			.plugin(tsify, {project: 'src/scripts/tsconfig.json'})
			.bundle();
		file.path = file.path.slice(0, -3) + '.js';
	}))
	.pipe(buffer())
	.pipe(sourcemaps.init({ loadMaps: true }))
	.pipe(uglify())
	.pipe(print(filepath => `built: ${filepath}`))
	.pipe(sourcemaps.write('.'))
	.pipe(gulp.dest('dist/public/js/'));
}

function compileServerTs () {
	return serverTs.src()
		.pipe(newer({dest: 'dist/', ext: '.js'}))
		.pipe(sourcemaps.init())
		.pipe(serverTs())
		.js.pipe(buffer())
		.pipe(uglify())
		.pipe(print(filepath => `built: ${filepath}`))
		.pipe(sourcemaps.write('.'))
		.pipe(gulp.dest('dist/'));
}

export const ts = gulp.parallel(compileClientTs, compileServerTs);

function copyAssets () {
	return gulp.src('src/public/**')
		.pipe(newer('dist/public/'))
		.pipe(print(filepath => `copied: ${filepath}`))
		.pipe(gulp.dest('dist/public/'));
}

function copyFavicons () {
	return gulp.src('src/favicon/public/**')
		.pipe(newer('dist/public/'))
		.pipe(print(filepath => `copied: ${filepath}`))
		.pipe(gulp.dest('dist/public/'));
}

function copyViews () {
	return gulp.src('src/views/**/*.pug')
		.pipe(newer('dist/views/'))
		.pipe(print(filepath => `copied: ${filepath}`))
		.pipe(gulp.dest('dist/views/'));
}

export const build = gulp.parallel(sass, ts, copyAssets, copyFavicons, copyViews);

function spawnServer () {
	return spawn('node', ['dist/app.js'], {
		env: process.env,
		stdio: 'inherit'
	});
}

function startServer (done) {
	spawnServer();
	done();
}

export function runServer (done) {
	const child = spawnServer();
	child.on('close', (code) => {
		done(`Server stopped with code ${code}.`);
	});
}

export const bRun = gulp.series(build, runServer);

function bsInit (done) {
	browserSync.init({
		proxy: 'localhost:' + Number(process.env.PORT),
		port: Number(process.env.PORT) + 1,
		open: false
	});
	done();
}

function reloadBs (done) {
	browserSync.reload();
	done();
}

function watch () {
	gulp.watch('src/scripts/**/*.ts', gulp.series(ts, reloadBs));
	gulp.watch('src/views/**/*.pug', gulp.series(copyViews, reloadBs));
	compileSass(gulpWatchSass('src/sass/**/*.scss')).pipe(browserSync.stream());
}

export const dev = gulp.series(build, startServer, bsInit, watch);

export function lintJs () {
	return gulp.src(['**/*.js', '!node_modules/**', '!dist/**'])
		.pipe(happiness())
		.pipe(happiness.format())
		.pipe(happiness.failAfterError());
}

function lintServerTs () {
	const program = tslint.Linter.createProgram('./src/tsconfig.json');
	return gulp.src(['./src/**/*.ts', '!src/scripts/**'])
		.pipe(gulpTslint({
			formatter: 'verbose',
			program: program,
			configuration: './tslint.json'
		}))
		.pipe(gulpTslint.report());
}

function lintClientTs () {
	const program = tslint.Linter.createProgram('./src/scripts/tsconfig.json');
	return gulp.src(['./src/scripts/**/*.ts'])
		.pipe(gulpTslint({
			formatter: 'verbose',
			program: program,
			configuration: './tslint.json'
		}))
		.pipe(gulpTslint.report());
}

export const lintTs = gulp.series(lintServerTs, lintClientTs);

export const lint = gulp.series(lintJs, lintTs);

export default build;
