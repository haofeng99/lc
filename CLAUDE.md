# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project overview

This is a LeetCode solutions practice repo in Go (module `lc`, Go 1.21). Solutions are organized by algorithm/data-structure category, with each `.go` file solving a specific LeetCode problem. The `main.go` file at the repo root is not the entry point for solutions — it contains Go language exploration snippets (slice/map initialization).

## Commands

```bash
# Run a single test file
go test ./sort/ -run TestSort -v

# Run all tests in a package
go test ./link_list/ -v

# Run a specific Go file (ad-hoc — most files are standalone packages)
go run ./sort/pancakeSort.go
```

There is no build system or CI. Most files are not runnable on their own (they contain only the solution function, no `main`). Test files exist only in a few packages (`sort/`, `link_list/`, `pointer/`).

## Architecture

Each top-level directory is a Go package named after a data structure or algorithm category:

| Directory | Focus |
|---|---|
| `sort/` | Sorting algorithms and problems solvable by sorting. `base_sort/` has foundational sorts (merge, quick, heap, bubble, shell, bucket, count). |
| `link_list/` | Singly and doubly linked list problems. `Node.go` / `DLLNode.go` define shared list node types. |
| `pointer/` | Two-pointer and sliding window problems (`slid_window/`). |
| `dynamic_programming/` | DP problems, further grouped by subcategory: `profit/`, `palindrome/`, `sub_arr_max_value/`, `base_pack/`, `dp_tree/`. |
| `tree/` | Tree problems (binary tree, BST). |
| `pre_sum/` | Prefix sum problems in 1D (`one_dim/`), 2D (`two_dim/`), and with index hashing (`pre_sum_idx_hash/`). |
| `search/` | Search algorithms: `binary_search/`, `bfs/`, `dfs/`. |
| `queue/` | Queue-based problems: `priority_queue/`, `monotonous_queue/`. |
| `stack/` | Stack-based problems: `monotonic_stack/`. |
| `hash/` | Hash-table problems (`index_hash/`). |
| `math/` | Math problems (`bit/` for bit manipulation). |
| `binary_indexed_tree/` | Fenwick tree (BIT) problems. |
| `divide_conquer/` | Divide-and-conquer problems. |
| `heap/` | Heap problems. |
| `enum/` | Enumeration problems. |
| `lc_weekend/` | Weekend contest problems. |

Shared types (e.g., `Node`, `ListNode`) are defined in the package that owns them. Some cross-package duplication exists (e.g., `Node` appears in both `link_list/` and `pointer/`).

The third-party dependency `github.com/erriapo/redblacktree` is declared in `go.mod` but not actively imported in any solution file.
