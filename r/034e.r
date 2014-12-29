#!/usr/bin/Rscript

require(compiler)

f <- factorial(0:9)

forEachDigit <- cmpfun(function(n, func) {
    div <- 1
    results <- NULL

    while (floor(n / div) > 0) {
        digit <- floor(n %% (div * 10) / div)
        results <- c(results, func(digit))
        div <- div * 10
    }

    results
})

fact <- cmpfun(function(x) { f[x+1] })

sum <- 0
for (n in 10:99999) {
    if (sum(forEachDigit(n, fact)) == n) {
        sum <- sum + n
    }
}

cat(sum, "\n")
