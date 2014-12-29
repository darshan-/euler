#!/usr/bin/Rscript

sumOfFifthPowOfDigits <- function(n) {
    sum <- 0
    div <- 1

    while (floor(n / div) > 0) {
        digit <- floor(n %% (div * 10) / div)
        sum <- sum + digit^5
        div <- div * 10
    }

    sum
}

sum <- 0
for (n in 10:295245) {
    if (sumOfFifthPowOfDigits(n) == n) {
        sum <- sum + n
        cat("Found:", n, "\n")
    }
}

cat(sum, "\n")
