package com.dalen.aoc

object Utils {
    fun getLines(file: String): List<String> {
        return this::class.java.classLoader.getResourceAsStream("$file.input")
            .bufferedReader()
            .readLines()
    }
}

data class Point(
    val x: Int,
    val y: Int
)

fun Point.move(direction: Direction): Point {
    return when (direction) {
        Direction.U -> Point(x, y - 1)
        Direction.D -> Point(x, y + 1)
        Direction.L -> Point(x - 1, y)
        Direction.R -> Point(x + 1, y)
    }
}

enum class Direction {
    U, D, L, R;
}