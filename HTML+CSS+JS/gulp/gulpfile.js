let project_folder = "dist";
let source_folder = "#src";

let path = {
    // Описание проекта для заказчика
    build: {
        html: project_folder + "/",
        css: project_folder + "/css/",
        js: project_folder + "/js/",
        img: project_folder + "/img/",
        fonts: project_folder + "/fonts/",
    },
    // Описание проекта для разработки
    src: {
        html: [source_folder + "/*.html", "!" + source_folder + "/_*.html"],
        css: source_folder + "/scss/style.scss",
        js: source_folder + "/js/script.js",
        img: source_folder + "/img/**/*.{jpg,png,svg,gif,ico,webp}",
        fonts: source_folder + "/fonts/*.ttf",
    },
    watch: {
        html: source_folder + "/**/*.html",
        css: source_folder + "/scss/**/*.scss",
        js: source_folder + "/js/**/*.js",
        img: source_folder + "/img/**/*.{jpg,png,svg,gif,ico,webp}",
    },
    clean: "./" + project_folder + "/"
}

let { src, dest } = require('gulp'),
    gulp = require('gulp'),                                 // Подключение Gulp
    browsersync = require("browser-sync").create(),         // Live server
    fileinclude = require("gulp-file-include"),             // Для сбора файла по частям в один файл
    del = require("del"),                                   // Для очистки проекта
    scss = require("gulp-sass"),                            // Предпроцессор SASS
    autoprefixer = require("gulp-autoprefixer"),            // Добавление префиксов в CSS
    group_media = require("gulp-group-css-media-queries"),  // Группирует медиафайлы в CSS
    clean_css = require("gulp-clean-css"),                  // Минимизирует CSS
    rename = require("gulp-rename"),                        // Переименовывает файлы
    uglify = require("gulp-uglify-es").default;

function browserSync(params) {
    browsersync.init({
        server: {
            baseDir: "./" + project_folder + "/"
        },
        port: 3000,
        notify: false
    })
}

function html() {
    return src(path.src.html)
        .pipe(fileinclude())
        .pipe(dest(path.build.html))
        .pipe(browsersync.stream())
}

function css() {
    return src(path.src.css)
        .pipe(
            scss({
                outputStyle: "expanded"
            })
        )
        .pipe(
            autoprefixer({
                overrideBrowserslist: ["last 5 versions"],
                cascade: true
            })
        )
        .pipe(
            group_media()
        )
        .pipe(dest(path.build.css)) // Выгрузка CSS файла
        .pipe(                      // Чистка CSS ("Минимизация")
            clean_css()
        )
        .pipe(                      // Переименовывание 
            rename({
                extname: ".min.css"
            })
        )
        .pipe(dest(path.build.css)) // Выгрузка переименнованного файла
        .pipe(browsersync.stream())
}

function js() {
    return src(path.src.js)
        .pipe(fileinclude())
        .pipe(dest(path.build.js))
        .pipe(
            uglify()
        )
        .pipe(
            rename({
                extname: ".min.js"
            })
        )
        .pipe(dest(path.build.js))
        .pipe(browsersync.stream())
}

function img() {
    return src(path.src.img)
        .pipe(fileinclude())
        .pipe(dest(path.build.html))
        .pipe(browsersync.stream())
}

function watchFiles(params) {              // Следим за файлами
    gulp.watch([path.watch.html], html);
    gulp.watch([path.watch.css], css);
    gulp.watch([path.watch.js], js);
}

function clean(params) {
    return del(path.clean);
}

let build = gulp.series(clean, gulp.parallel(js, css, html)); // Процесс сборки и выполнения
let watch = gulp.parallel(build, watchFiles, browserSync);    // Серверный процесс 

exports.js = js;
exports.css = css;
exports.html = html;
exports.build = build;
exports.watch = watch;
exports.default = watch;
