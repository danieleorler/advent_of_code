package com.dalen.aoc

import com.dalen.aoc.Utils.getLines

class Day08 {

    private val forest: MutableList<List<Int>> = mutableListOf()
    val visible: MutableMap<String, MutableSet<String>> = mutableMapOf()

    fun mapForest() {
        getLines("08").asSequence()
            .map { it.toList() }
            .map { it.map { c -> c.toString().toInt()} }
            .forEach { forest.add(it) }
    }

    fun findVisible() {
        for (i in forest.indices) {
            findVisibleInRow(i)
        }

        for (i in forest[0].indices) {
            findVisibleInColumn(i)
        }
    }

    fun findBestTree(): Int {
        var best = 0
        for (row in forest.indices) {
            for (column in forest[0].indices) {
                val score = scenicScore(row, column)
                if (score > best) {
                    best = score
                }
            }
        }

        return best
    }


    private fun findVisibleInRow(row: Int) {
        var max = -1
        for (i in forest[row].indices) {
            if (forest[row][i] > max) {
                max = forest[row][i]
                addVisibleFrom("$row,$i", "L")
            }
        }

        max = -1
        for (i in (forest[row].size - 1 downTo   0)) {
            if (forest[row][i] > max) {
                max = forest[row][i]
                addVisibleFrom("$row,$i", "R")
            }
        }
    }

    private fun findVisibleInColumn(column: Int) {
        var max = -1
        for (i in forest.indices) {
            if (forest[i][column] > max) {
                max = forest[i][column]
                addVisibleFrom("$i,$column", "T")
            }
        }

        max = -1
        for (i in (forest.size - 1 downTo   0)) {
            if (forest[i][column] > max) {
                max = forest[i][column]
                addVisibleFrom("$i,$column", "D")
            }
        }
    }

    private fun addVisibleFrom(tree: String, direction: String) {
        if (visible.containsKey(tree)) {
            visible[tree]!!.add(direction)
        } else {
            visible[tree] = mutableSetOf(direction)
        }
    }

    private fun scenicScore(row: Int, column: Int): Int {
        val tree = forest[row][column]
        if (row == 0 || column == 0) {
            return 0
        }

        if (row == forest.size - 1 || column == forest[0].size - 1) {
            return 0
        }

        var left = 0
        var nextColumn = column - 1
        var reached = false
        while (nextColumn >= 0 && nextColumn < forest[0].size && !reached) {
            left += 1
            reached = forest[row][nextColumn] >= tree
            nextColumn--
        }

        var right = 0
        nextColumn = column + 1
        reached = false
        while (nextColumn >= 0 && nextColumn < forest[0].size && !reached) {
            right += 1
            reached = forest[row][nextColumn] >= tree
            nextColumn++
        }

        var top = 0
        var nextRow = row - 1
        reached = false
        while (nextRow >= 0 && nextRow < forest.size && !reached) {
            top += 1
            reached = forest[nextRow][column] >= tree
            nextRow--
        }

        var bottom = 0
        nextRow = row + 1
        reached = false
        while (nextRow >= 0 && nextRow < forest.size && !reached) {
            bottom += 1
            reached = forest[nextRow][column] >= tree
            nextRow++
        }

        return left * right * top * bottom
    }

}

fun main() {
    val day = Day08()
    day.mapForest()
    day.findVisible()

    println("Solution1: ${day.visible.keys.count()}")
    println("Solution2: ${day.findBestTree()}")
}