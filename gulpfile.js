'use strict';

var Fiber = require('fibers');
var gulp = require('gulp');
var sass = require('gulp-sass');

sass.compiler = require('sass');

gulp.task('sass', function () {
  return gulp.src('./public_src/styles/*.scss')
    .pipe(sass({fiber: Fiber}).on('error', sass.logError))
    .pipe(gulp.dest('./public/styles'));
});

gulp.task('sass:watch', function () {
  gulp.watch('./public_src/styles/*.scss', gulp.series('sass'));
});