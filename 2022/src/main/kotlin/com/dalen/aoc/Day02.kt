package com.dalen.aoc

import com.dalen.aoc.Shape.Companion.fromLetter
import com.dalen.aoc.Utils.getLines

enum class Shape(private val letters: List<Char>, val points: Int, private val w: Char, private val l: Char) {
    ROCK(listOf('A', 'X'), 1, 'C', 'B'),
    PAPER(listOf('B', 'Y'), 2, 'A', 'C'),
    SCISSOR(listOf('C', 'Z'), 3, 'B', 'A');

    fun winsAgainst(): Shape {
        return fromLetter(this.w)
    }

    fun losesAgainst(): Shape {
        return fromLetter(this.l)
    }

    fun vs(other: Shape): Int {
        if (this == other) {
            return 3
        }
        if (this.losesAgainst() == other) {
            return 0
        }
        return 6
    }

    companion object {
        fun fromLetter(letter: Char): Shape {
            return values().first { it.letters.contains(letter) }
        }
    }
}

class Day02 {

    fun solve1(): Int {
        return getLines("02.input")
            .map { listOf(it[0], it[2]) }
            .map { listOf(Shape.fromLetter(it[0]), Shape.fromLetter(it[1])) }
            .sumOf { it[1].vs(it[0]) + it[1].points }
    }

    fun solve2(): Int {

        val plays = mapOf<Char, (shape: Shape) -> Shape>(
            'X' to { s -> s.winsAgainst() },
            'Y' to { s -> s },
            'Z' to { s -> s.losesAgainst() }
        )

        return getLines("02.input")
            .asSequence()
            .map { listOf(it[0], it[2]) }
            .map { listOf(fromLetter(it[0]), plays[it[1]]!!(fromLetter(it[0]))) }
            .sumOf { it[1].vs(it[0]) + it[1].points }
    }
}

fun main() {
    val day2 = Day02()

    println("Solution 1: ${day2.solve1()}")
    println("Solution 2: ${day2.solve2()}")
}