#!/usr/bin/env rake

ROOT_DIR = File.dirname(__FILE__)

INSTALL_DIR = ENV["INSTALL_DIR"] || File.join(ROOT_DIR, "build")
EXECUTABLE  = File.join(INSTALL_DIR, "ppl")

task :default => :build

task :build do
  Dir.chdir(ROOT_DIR) do
    sh "go build -o #{EXECUTABLE} src/ppl.go"
  end
end
