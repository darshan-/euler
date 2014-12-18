#!/usr/bin/Rscript

limit <- NULL
nums <- NULL

setup <- function(lim) {
    limit <<- lim
    nums <<- rep(NA, limit)
    nums[1] <- FALSE
}

INITIAL.LIMIT <- 1e5
GOAL <- 10001
setup(INITIAL.LIMIT)

repeat {
    for (num in 2:limit) {
        if (!is.na(nums[num])) { next }
        nums[num] <- TRUE
        if (num < limit / 2) {
            nums[seq(num*2, limit, num)] <- FALSE
        }
    }

    if (!is.na(which(nums)[GOAL])) { break }

    setup(limit * 2)
}

cat(which(nums)[GOAL], "\n")
