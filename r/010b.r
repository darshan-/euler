#!/usr/bin/Rscript

LIMIT <<- 2e6
nums <<- rep(NA, LIMIT)
nums[1] <- FALSE
sum <- 0

for (num in 2:LIMIT) {
    if (!is.na(nums[num])) { next }

    sum <- sum + num

    if (num < LIMIT / 2) {
        nums[seq(num*2, LIMIT, num)] <- FALSE
    }
}

cat(sum, "\n")
