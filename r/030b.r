#!/usr/bin/Rscript

require(compiler)

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

pow5 <- cmpfun(function(x) { x^5 })

sum <- 0
for (n in 10:295245) {
    if (sum(forEachDigit(n, pow5)) == n) {
        sum <- sum + n
        cat("Found:", n, "\n")
    }
}

cat(sum, "\n")
