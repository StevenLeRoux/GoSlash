module.exports = function (grunt) {

  // Project configuration.
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    copy: {
      build: {
        files: [{
          expand: true,
          cwd: 'src',
          src: '**/*.html',
          dest: 'build'
        }]
      },
    },
    uglify: {
      build: {
        options: {},
        files: [{
          expand: true,
          cwd: 'src',
          src: '**/*.js',
          dest: 'build'
        }]
      }
    },
    less: {
      build: {
        options: {},
        files: [{
          expand: true,
          cwd: 'src',
          src: '**/*.less',
          dest: 'build'
        }]
      }
    },
    postcss: {
      options: {
        map: true,
        processors: [
          require('autoprefixer')()
        ]
      },
      build: {
        src: 'build/**/*.less'
      }
    },
    processhtml: {
      options: {},
      dist: {
        files: [{
          expand: true,
          cwd: 'build',
          src: '**/*.html',
          dest: 'dist'
        }]
      }
    },
    clean: {
      build: ['build'],
      dist: ['dist'],
    },
    watch: {
      dev: {
        files: ['src/**/*'],
        tasks: ['dist'],
        options: {
          atBegin: true,
        },
      }
    }
  });

  grunt.loadNpmTasks('grunt-contrib-clean');
  grunt.loadNpmTasks('grunt-contrib-less');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-postcss');
  grunt.loadNpmTasks('grunt-processhtml');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-copy');

  grunt.registerTask('default', ['clean:dist', 'dist']);
  grunt.registerTask('dist', ['copy:build', 'uglify:build', 'less:build', 'postcss:build', 'processhtml:dist']);
  grunt.registerTask('dev', ['watch:dev']);
};
