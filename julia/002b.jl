LIMIT = 4*1000*1000
sum = 0

nextfib = (() ->
    begin
        a, b = 0, 1

        function()
            a, b = b, a + b
            b
        end
    end)()

while true
    f = nextfib()
    if f > LIMIT
        break
    end

    if f % 2 == 0
        sum += f
    end
end

println(sum)
