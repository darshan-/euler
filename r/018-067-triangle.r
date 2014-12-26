triangle <- function(file) {
    con  <- file(file, open = "r")
    nums <- list()

    while (length(oneLine <- readLines(con, n = 1, warn = FALSE)) > 0) {
        v <- list(as.numeric(strsplit(oneLine, " ")[[1]]))
        nums <- c(nums, v)
    }

    for (r in (length(nums)-1):1) {
        for (c in 1:length(nums[[r]])) {
            nums[[r]][c] <- nums[[r]][c] + max(nums[[r+1]][c], nums[[r+1]][c+1])
        }
    }

    nums[[1]][1]
}
