#!/usr/bin/Rscript

palindromeOf <- function (n) {
    digits <- NULL
    dup <- n

    while (dup > 0) {
        digits <- append(digits, dup %% 10)
        dup <- floor(dup / 10)
    }

    for (digit in digits) {
        n <- n * 10
        n <- n + digit
    }

    n
}

largestNDigitNum <- function(n) {
    10^n - 1
}

smallestNDigitNum <- function(n) {
    10^(n-1)
}

isProductOfNDigits <- function(num, nDigits) {
    max <- largestNDigitNum(nDigits)
    min <- smallestNDigitNum(nDigits)

    i <- max
    while (i > min) {
        if (num %% i == 0 && num / i < max) {
            return(TRUE)
        }

        i <- i - 1
    }

    FALSE
}

largestPalOfNDigits <- function(nDigits) {
    i <- largestNDigitNum(nDigits)
    while (i > 0) {
        pal <- palindromeOf(i)
        if (isProductOfNDigits(pal, nDigits)) {
            return(pal)
        }
        i <- i-1
    }
}

cat(largestPalOfNDigits(3), "\n")
