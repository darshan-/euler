#!/usr/bin/Rscript

f <- factorial(0:9)

require(compiler)
sumOfFactorialOfDigits <- cmpfun(function(n) {
    sum <- 0
    div <- 1

    while (floor(n / div) > 0) {
        digit <- floor(n %% (div * 10) / div)
        sum <- sum + f[digit+1]
        div <- div * 10
    }

    sum
})

sum <- 0
for (n in 10:99999) { ## Question excludes single-digit numbers, and I *think* five digits is max possible
    if (sumOfFactorialOfDigits(n) == n) {
        sum <- sum + n
    }
}

cat(sum, "\n")
