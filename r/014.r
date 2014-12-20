#!/usr/bin/Rscript

GOAL <- 1e6

max   <- 1
which <- 1

for (i in 1:GOAL) {
    count <- 1
    j <- i
    while (j > 1) {
        count <- count + 1

        if (j %% 2 == 0) {
            j <- j / 2
        } else {
            j <- 3 * j + 1
        }
    }
    if (count > max) {
        max <- count
        which <- i
        ##cat("Max:", max, "starts at", which, "\n")
    }
}

cat(which, "\n")
