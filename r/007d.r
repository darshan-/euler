#!/usr/bin/Rscript

LIMIT <- 100000
GOAL <- 10001
nums <- rep(NA, LIMIT)
nums[1] <- FALSE

for (num in 2:LIMIT) {
    if (!is.na(nums[num])) { next }
    nums[num] <- TRUE
    if (num < LIMIT / 2) {
        nums[seq(num*2, LIMIT, num)] <- FALSE
    }
}

cat(which(nums)[GOAL], "\n")
