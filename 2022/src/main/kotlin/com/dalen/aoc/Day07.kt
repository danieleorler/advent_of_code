package com.dalen.aoc

import com.dalen.aoc.Utils.getLines

class Day07 {

    private val currentPath = ArrayDeque<String>()
    private val fs = mutableMapOf<String, Int>()
    private val fileRegex = """^(\d+)\s.*""".toRegex()

    init {
        fs()
    }

    private fun fs() {
        getLines("07")
            .forEach {
                when {
                    it.startsWith("$ cd") -> cd(it.substring("$ cd ".length, it.length))
                    it == "$ ls" -> Unit
                    !it.startsWith("$") -> sum(it)
                }
            }
    }

    private fun cd(next: String) {
        if (next == "..") {
            currentPath.removeLast()
        } else {
            currentPath.addLast(next)
        }
    }

    private fun sum(file: String) {
        if (!file.startsWith("dir")) {
            val (size) = fileRegex.find(file)!!.destructured
            var path = ""
            for (i in currentPath.indices) {
                path += "${currentPath[i]}/"
                fs[path] = (fs[path] ?: 0) + size.toInt()
            }
        }
    }

    fun sol1(): Int {
        return fs.values.filter { it < 100000  }.sum()
    }

    fun sol2(): Int {
        val total = 70000000
        val required = 30000000
        val available = total - fs["//"]!!
        val toFree = required - available
        return fs.values.filter { it >= toFree }.min()
    }

}

fun main() {
    val day = Day07()

    println("Solution1: ${day.sol1()}")
    println("Solution2: ${day.sol2()}")
}