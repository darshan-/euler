#!/usr/bin/Rscript

LIMIT <<- 2e6
SQRT <<- floor(sqrt(LIMIT))
nums <<- vector(length=LIMIT)
sum <- 0

for (num in 2:SQRT) {
    if (nums[num]) { next }
    sum <- sum + num
    nums[seq(num^2, LIMIT, num)] <- TRUE
}

for (num in (SQRT + 1):LIMIT) {
    if (!nums[num]) {
        sum <- sum + num
    }
}

cat(sum, "\n")
