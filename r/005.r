#!/usr/bin/Rscript

n <- 1
RANGE <- 1:20

while (TRUE) {
    if (all(n %% RANGE == 0)) { break }
    n <- n + 1
}

cat(n, "\n")
