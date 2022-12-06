package com.dalen.aoc

import com.dalen.aoc.Utils.getLines

class Day06 {

    private fun findMarker(signal: String, length: Int): Int {
        var start = 0
        var end = length
        var marker = signal.substring(start, end)

        while (!marker.isMarker(length)) {
            start += 1
            end += 1
            marker = signal.substring(start, end)
        }
        return end
    }

    private fun String.isMarker(length: Int): Boolean {
        val str = this.toCharArray().sorted()
        var allDifferent = true
        for (i in 0 until str.size - 1) {
            allDifferent = allDifferent && (str[i] < str[i+1])
        }
        return this.length == length && allDifferent
    }

    fun solve1(): Int = findMarker(getLines("06").first(), 4)
    fun solve2(): Int = findMarker(getLines("06").first(), 14)




}

fun main() {
    val day = Day06()

    println("solution 1: ${day.solve1()}")
    println("solution 2: ${day.solve2()}")
}