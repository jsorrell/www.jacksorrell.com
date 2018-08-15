'use strict';

require('dotenv').config();

var gulp = require('gulp');
var autoprefixer = require('autoprefixer');
var bs = require('browser-sync').create();
var minifycss = require('gulp-clean-css');
var postcss = require('gulp-postcss');
var sass = require('gulp-sass');
var spawn = require('child_process').spawn;
var tsProject = require('gulp-typescript').createProject('scripts/tsconfig.json');
var uglify = require('gulp-uglify');
var buffer = require('vinyl-buffer');

gulp.task('sass', function () {
	return gulp.src('sass/**/*.scss')
		.pipe(sass().on('error', sass.logError))
		.pipe(postcss([autoprefixer({
			cascade: false
		})]))
		.pipe(minifycss())
		.pipe(gulp.dest('public/css/'))
		.pipe(bs.stream());
});

gulp.task('ts', function () {
	var ret = tsProject.src()
		.pipe(tsProject())
		.js.pipe(buffer())
		.pipe(uglify())
		.pipe(gulp.dest('public/js/'));
	bs.reload();
	return ret;
});

gulp.task('build', gulp.parallel('sass', 'ts'));
gulp.task('default', gulp.series('build'));

gulp.task('run', function (done) {
	const child = spawn('node', ['bin/www'], {
		env: process.env,
		stdio: 'inherit'
	});
	child.on('close', (code) => {
		done(`Server stopped with code ${code}.`);
	});
});

gulp.task('bs', gulp.series('build', gulp.parallel('run', function () {
	gulp.watch('sass/**/*.scss', gulp.series('sass'));
	gulp.watch('scripts/**/*.ts', gulp.series('ts'));
	gulp.watch('views/**/*.pug', bs.reload);

	bs.init({
		proxy: 'localhost:' + Number(process.env.PORT),
		port: Number(process.env.PORT) + 1,
		open: false
	});
})));
