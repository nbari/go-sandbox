https://blog.alexellis.io/prometheus-monitoring/

You can generate a hash like this:

    $ curl localhost:8080/hash -d "my_input_value_here"
    49b8c4256d603a68ee9bcd95f8e11eed784189bd40c354950014b6d7f7263d6c


This will now show up if we curl localhost:8080/metrics:

    # HELP hash_seconds Time taken to create hashes
    # TYPE hash_seconds histogram
    hash_seconds_bucket{code="200",le="1"} 2
    hash_seconds_bucket{code="200",le="2.5"} 2
    hash_seconds_bucket{code="200",le="5"} 2
    hash_seconds_bucket{code="200",le="10"} 2
    hash_seconds_bucket{code="200",le="+Inf"} 2
    hash_seconds_sum{code="200"} 9.370800000000002e-05
    hash_seconds_count{code="200"} 2
