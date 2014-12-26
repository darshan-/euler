triangle <- function(filename) {
    con  <- file(filename, open = "r")
    nums <- list()

    while (length(line <- readLines(con, n = 1, warn = FALSE)) > 0) {
        nums <- c(nums, list(as.numeric(strsplit(line, " ")[[1]])))
    }

    for (r in (length(nums)-1):1) {
        for (c in 1:length(nums[[r]])) {
            nums[[r]][c] <- nums[[r]][c] + max(nums[[r+1]][c], nums[[r+1]][c+1])
        }
    }

    nums[[1]][1]
}
