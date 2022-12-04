package com.dalen.aoc

import com.dalen.aoc.Utils.getLines

class Day04 {

    private fun List<String>.toRanges(): List<List<Pair<Int, Int>>> {
        return this
            .map { it.split(",") }
            .map { listOf(it[0].split("-"), it[1].split("-")) }
            .map { listOf(
                Pair(it[0][0].toInt(), it[0][1].toInt()),
                Pair(it[1][0].toInt(), it[1][1].toInt()),
            )}
    }

    fun solve1(): Int {
        return getLines("04")
            .toRanges()
            .count { it[0].contains(it[1]) || it[1].contains(it[0]) }
    }

    fun solve2(): Int {
        return getLines("04")
            .toRanges()
            .count { it[0].overlap(it[1]) }
    }

    private fun Pair<Int, Int>.contains(other: Pair<Int, Int>): Boolean {
        return this.first <= other.first && this.second >= other.second
    }

    private fun Pair<Int, Int>.overlap(other: Pair<Int, Int>): Boolean {
        return this.first >= other.first && this.first <= other.second
                || this.second >= other.first && this.first <= other.second
    }
}

fun main() {
    val day = Day04()
    println("Solution1: ${day.solve1()}")
    println("Solution1: ${day.solve2()}")
}