package com.dalen.aoc

import com.dalen.aoc.Utils.getLines

class Day05 {
    private val instructionRegex = """^move\s(\d+)\sfrom\s(\d+)\sto\s(\d+)$""".toRegex()

    private fun Sequence<String>.toInstructions(): Sequence<List<Int>> {
        return this
            .map { instructionRegex.find(it) }
            .map { val (a,b,c) = it!!.destructured; listOf(a.toInt(), b.toInt(), c.toInt()) }
    }

    fun solve1(stacks: Map<Int, ArrayDeque<String>>): String {
        getLines("05").asSequence()
            .toInstructions()
            .forEach {
                for (i in (0 until it[0])) {
                    stacks[it[2]]!!.addFirst(stacks[it[1]]!!.removeFirst())
                }
            }
        return stacks.map { (_, v) -> v.first() }.joinToString("")
    }

    fun solve2(stacks: Map<Int, ArrayDeque<String>>): String {
        getLines("05").asSequence()
            .toInstructions()
            .forEach {
                val toMove = ArrayDeque<String>()
                for (i in (0 until it[0])) {
                    toMove.addFirst(stacks[it[1]]!!.removeFirst())
                }
                while (!toMove.isEmpty())
                {
                    stacks[it[2]]!!.addFirst(toMove.removeFirst())
                }

            }
        return stacks.map { (_, v) -> v.first() }.joinToString("")
    }
}

fun main() {
    val day = Day05()
    println("Solution1: ${day.solve1(initialStacks())}")
    println("Solution2: ${day.solve2(initialStacks())}")
}

private fun initialStacks(): Map<Int, ArrayDeque<String>> {
    return mapOf(
        1 to ArrayDeque(listOf("M", "S", "J", "L", "V", "F", "N", "R")),
        2 to ArrayDeque(listOf("H", "W", "J", "F", "Z", "D", "N", "P")),
        3 to ArrayDeque(listOf("G", "D", "C", "R", "W")),
        4 to ArrayDeque(listOf("S", "B", "N")),
        5 to ArrayDeque(listOf("N", "F", "B", "C", "P", "W", "Z", "M")),
        6 to ArrayDeque(listOf("W", "M", "R", "P")),
        7 to ArrayDeque(listOf("W", "S", "L", "G", "N", "T", "R")),
        8 to ArrayDeque(listOf("V", "B", "N", "F", "H", "T", "Q")),
        9 to ArrayDeque(listOf("F", "N", "Z", "H", "M", "L")),
    )
}