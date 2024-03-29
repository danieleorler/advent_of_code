package com.dalen.aoc

import com.dalen.aoc.Utils.getLines

class Day01 {
    fun process(): List<Int> {
        val calories = mutableListOf<Int>()
        var sum = 0
        getLines("01")
            .forEach {
                if (it.isBlank()) {
                    calories.add(sum)
                    sum = 0
                } else {
                    sum += it.toInt()
                }
            }

        calories.sort()
        return calories
    }
}

fun main() {
    val day = Day01()

    val calories = day.process()

    println("Solution 1: ${calories.last()}")
    println("Solution 2: ${calories.takeLast(3).sum()}")

}
