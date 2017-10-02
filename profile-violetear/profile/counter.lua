--  wrk -c2 -d1s -t1 -s counter.lua http://0:8080

counter = 0

request = function()
    if counter % 100 == 0 then
	    path = "/"
    else
	    path = path .. counter .. "/"
    end
    wrk.headers["X-Counter"] = counter
    counter = counter + 1
    return wrk.format(nil, path)
end
