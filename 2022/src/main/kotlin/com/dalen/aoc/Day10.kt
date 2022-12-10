package com.dalen.aoc

import com.dalen.aoc.Utils.getLines

class Day10 {
    fun solve1(): Int {

        var reg = 1
        val cycles = mutableListOf<Int>()
        cycles.add(reg)


        getLines("10")
            .map {
                if (it == "noop") {
                    cycles.add(reg)
                } else {
                    val toAdd = it.split(" ")[1].toInt()
                    cycles.add(reg)
                    reg += toAdd
                    cycles.add(reg)
                }
            }

        return listOf(20, 60, 100, 140, 180, 220)
            .sumOf { it * cycles[it-1] }
    }

    fun solve2() {
        val screen = initScreen()
        var sprite = 1
        var crt = 0

        getLines("10")
            .forEach {
                if (it == "noop") {
                    drawPixel(crt++, sprite, screen)
                } else {
                    drawPixel(crt++, sprite, screen)
                    drawPixel(crt++, sprite, screen)
                    val toAdd = it.split(" ")[1].toInt()
                    sprite += toAdd
                }
            }

        printScreen(screen)
    }

    private fun drawPixel(crt: Int, sprite: Int, screen: MutableList<MutableList<String>>) {
        if ((crt%40).isIn(sprite%40)) {
            screen[crt/40][crt%40] = "#"
        }
    }

    private fun Int.isIn(sprite: Int): Boolean {
        return this >= sprite - 1 && this <= sprite + 1
    }

    private fun initScreen(): MutableList<MutableList<String>> {
        val screen = mutableListOf<MutableList<String>>()
        for (i in (0 until 6)) {
            val row = mutableListOf<String>()
            for (j in (0 until 40)) {
                row.add(" ")
            }
            screen.add(row)
        }
        return screen
    }

    private fun printScreen(screen: MutableList<MutableList<String>>) {
        for (row in screen) {
            for (cell in row) {
                print(cell)
            }
            println()
        }
    }
}

fun main() {
    val day = Day10()

    println("Solution1: ${day.solve1()}")
    println("Solution2: ${day.solve2()}")
}