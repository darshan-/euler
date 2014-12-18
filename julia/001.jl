function col(dict, step, limit)
    i = step
    while i < limit
        n = i
        d[n] = n
        i += step
    end
end

LIMIT = 1000

d = Dict{Int, Int}()

col(d, 3, LIMIT)
col(d, 5, LIMIT)

println(sum(keys(d)))

