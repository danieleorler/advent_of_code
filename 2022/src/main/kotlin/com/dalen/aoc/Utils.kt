package com.dalen.aoc

import kotlin.math.abs

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
        Direction.S -> this
    }
}

fun Point.distanceTo(to: Point): Int {
    return abs(x - to.x) + abs(y - to.y)
}

fun Point.directionTo(to: Point): Direction {
    if (x == to.x && y < to.y)
        return Direction.D
    if (x == to.x && y > to.y)
        return Direction.U
    if (x < to.x && y == to.y)
        return Direction.R
    if (x > to.x && y == to.y)
        return Direction.L
    return Direction.S
}

enum class Direction {
    U, D, L, R, S;
}