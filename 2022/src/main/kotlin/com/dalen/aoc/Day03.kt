package com.dalen.aoc

import com.dalen.aoc.Utils.getLines

class Day03 {
    fun solve1(): Int {
        return getLines("03")
            .asSequence()
            .map {
                listOf(
                    HashSet(it.substring(0 until it.length / 2).toList()),
                    HashSet(it.substring(it.length / 2 until it.length).toList())
                )
            }
            .map { it[0].intersect(it[1]).first() }
            .map { it.code }
            .map { if (it > 90) it - 96 else it - 64 + 26 }
            .sum()
    }

    fun solve2(): Int {
        return getLines("03")
            .asSequence()
            .map { it.toSet() }
            .chunked(3)
            .map { g -> g[0].intersect(g[1]).intersect(g[2]).first() }
            .map { it.code }
            .map { if (it > 90) it - 96 else it - 64 + 26 }
            .sum()
    }
}

fun main() {
    val day = Day03()

    println("Solution 1: ${day.solve1()}")
    println("Solution 2: ${day.solve2()}")
}