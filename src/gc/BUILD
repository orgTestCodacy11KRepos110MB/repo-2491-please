go_library(
    name = "gc",
    srcs = ["gc.go"],
    visibility = ["PUBLIC"],
    deps = [
        "//src/cli",
        "//src/cli/logging",
        "//src/core",
        "//src/fs",
        "//src/parse/asp",
        "//src/scm",
    ],
)

go_test(
    name = "gc_test",
    srcs = ["gc_test.go"],
    deps = [
        ":gc",
        "//third_party/go:testify",
    ],
)

go_test(
    name = "rewrite_test",
    srcs = ["rewrite_test.go"],
    data = [
        "test_data",
    ],
    deps = [
        ":gc",
        "//third_party/go:testify",
    ],
)
