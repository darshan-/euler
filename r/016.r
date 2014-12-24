#!/usr/bin/Rscript

library(gmp)

s <- format(as.bigz(2^1000))
sum <- 0

for (i in 1:nchar(s)) {
    sum <- sum + as.integer(substr(s, i, i))
}

cat(sum, "\n")
