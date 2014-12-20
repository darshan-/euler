#!/usr/bin/Rscript

countFactors <- function(n) {
    count <- 0
    i <- 0

    while (i <- i + 1) {
        if (n %% i == 0) { # i is a factor
            count <- count + 1
            if (i * i == n) {
                return(count * 2 - 1)
            } else if (i * i > n) {
                return((count - 1) * 2)
            }
        }
        if (i * i > n) {
            return(count * 2)
        }
    }
}

GOAL <- 500

triNum <- 0
i <- 0
while(i <- i + 1) {
    triNum <- triNum + i
    if (countFactors(triNum) > GOAL) { break }
}

cat(triNum, "\n")
