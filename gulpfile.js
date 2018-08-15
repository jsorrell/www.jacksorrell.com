'use strict';

require('dotenv').config();

var gulp = require('gulp');
var autoprefixer = require('autoprefixer');
var bs = require('browser-sync').create();
var minifycss = require('gulp-clean-css');
// var minifyhtml = require('gulp-htmlmin');
var postcss = require('gulp-postcss');
// var pug = require('gulp-pug');
var sass = require('gulp-sass');


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

gulp.task('build', gulp.parallel('sass'));
gulp.task('default', gulp.series('build'));

gulp.task('bs', gulp.series('build', function () {
	gulp.watch('sass/**/*.scss', ['sass']);
	// gulp.watch(tsFiles, ['ts']);

	bs.init({
		proxy: 'localhost:' + Number(process.env.PORT),
		port: Number(process.env.PORT) + 1,
		open: false
	});
}));
