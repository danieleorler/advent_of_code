package com.dalen.aoc

import com.dalen.aoc.Utils.getLines
import com.fasterxml.jackson.databind.JsonNode
import com.fasterxml.jackson.databind.ObjectMapper
import java.lang.Integer.min
import kotlin.math.sign

class Day13 {

    private val mapper = ObjectMapper()

    fun solve1(): Int {
        val res = getLines("13")
            .chunked(3)
            .map { it.take(2) }
            .map { compare(mapper.readTree(it[0]), mapper.readTree(it[1])) }

        var sum = 0
        for (i in res.indices) {
            if (res[i] < 0) {
                sum += i + 1
            }
        }

        return sum
    }

    fun solve2(): Int {
        val input = getLines("13").toMutableList()
        val dividers = listOf("[[2]]", "[[6]]")
        input.addAll(dividers)
        val res = input
            .asSequence()
            .filter { it.isNotEmpty() }
            .map { mapper.readTree(it) }
            .sortedWith { a, b -> compare(a, b) }
            .toList()

        var result = 1
        for (i in res.indices) {
            if (dividers.contains("${res[i]}")) {
                result *= i + 1
            }
        }
        return result
    }

    private fun compare(left: JsonNode, right: JsonNode): Int {
        if (left.isArray && !right.isArray) {
            return compare(left, mapper.createArrayNode().add(right.intValue()))
        }
        if (!left.isArray && right.isArray) {
            return compare(mapper.createArrayNode().add(left.intValue()), right)
        }
        if (left.isArray && right.isArray) {
            var res = 0
            var i = 0
            while (res == 0 && i < min(left.size(), right.size())) {
                res += compare(left.get(i), right.get(i))
                i++
            }
            if (res != 0) {
                return res
            }
            if (left.size() > right.size()) {
                return 1
            }
            if (left.size() < right.size()) {
                return -1
            }
            return 0
        }

        return (left.intValue() - right.intValue()).sign
    }
}

fun main() {
    val day = Day13()

    println("Solution1: ${day.solve1()}")
    println("Solution2: ${day.solve2()}")
}
