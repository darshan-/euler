#!/usr/bin/Rscript

library(gmp)

s <- format(factorialZ(100))
sum <- 0

for (i in 1:nchar(s)) {
    sum <- sum + as.integer(substr(s, i, i))
}

cat(sum, "\n")
