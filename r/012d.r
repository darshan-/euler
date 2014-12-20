#!/usr/bin/Rscript

dyn.load("012d.so")

GOAL <- 500

ans <- .C("firstWithNFactors", as.integer(GOAL), as.integer(c(0)))[[2]]

cat(ans, "\n")
