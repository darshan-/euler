#!/usr/bin/Rscript

require(compiler)
enableJIT(3)

f <- factorial(0:9)

sum <- 0
for (n in 10:99999) { ## Question excludes single-digit numbers, and I *think* five digits is max possible
    sum2 <- 0
    div <- 1

    while (floor(n / div) > 0) {
        digit <- floor(n %% (div * 10) / div)
        sum2 <- sum2 + f[digit+1]
        div <- div * 10
    }

    if (sum2 == n) {
        sum <- sum + n
    }
}

cat(sum, "\n")
