package com.dalen.aoc

import com.dalen.aoc.Utils.getLines
import java.lang.Exception

class Day14 {

    fun solve1(): Int {
        val scan = getLines("14")
            .map { it.split(" -> ") }
            .map { it.map { p -> p.split(",") }.map { c -> Point(x = c[0].toInt(), y = c[1].toInt()) } }

        val map = scan.toMap()

        val (tl, _) = scan.boundaries()
        var done = false
        while (!done) {
            try {
                map.update(Point(500, 0).fallIn(map, tl.x))
            } catch (e: Exception) {
                done = true
            }
        }

        return map.count('o')
    }

    fun solve2(): Int {
        val scan = getLines("14")
            .map { it.split(" -> ") }
            .map { it.map { p -> p.split(",") }.map { c -> Point(x = c[0].toInt(), y = c[1].toInt()) } }
            .toMutableList()

        val (tl, br) = scan.boundaries()

        scan.add(listOf(Point(x = tl.x - 200, y = br.y + 2), Point(x = br.x + 200, y = br.y + 2)))

        val (ntl, _) = scan.boundaries()

        val map = scan.toMap()

        var last = Point(0, 0)
        while (last != Point(500 - ntl.x, 0)) {
            try {
                last = Point(500, 0).fallIn(map, ntl.x)
                map.update(last)
            } catch (e: Exception) {
                map.print()
                println(e.message)
                throw e
            }
        }

        return map.count('o')
    }

    private fun List<List<Point>>.boundaries(): Pair<Point, Point> {
        val scan = this.toMutableList()
        scan.add(listOf(Point(500, 0)))

        var minx = Int.MAX_VALUE
        var maxx = Int.MIN_VALUE
        var miny = Int.MAX_VALUE
        var maxy = Int.MIN_VALUE
        for (row in scan) {
            for (point in row) {
                with(point) {
                    if (x < minx)
                        minx = x
                    if (x > maxx)
                        maxx = x
                    if (y < miny)
                        miny = y
                    if (y > maxy)
                        maxy = y
                }
            }
        }

        val min = Point(x = minx, y = miny)
        val max = Point(x = maxx, y = maxy)

        return Pair(min, max)
    }

    private fun List<List<Point>>.toMap(): MutableList<MutableList<Char>> {
        val (tl, br) = this.boundaries()
        val minx = tl.x
        val map = mutableListOf<MutableList<Char>>()
        for (y in (0..br.y)) {
            val row = mutableListOf<Char>()
            for (x in (0..br.x - minx)) {
                row.add('.')
            }
            map.add(row)
        }

        for (rock in this.unwrap()) {
            map[rock.y][rock.x - minx] = '#'
        }

        return map
    }

    private fun List<List<Point>>.unwrap(): Set<Point> {
        val rocks = mutableSetOf<Point>()
        for (structure in this) {
            for (i in (0 until structure.size - 1)) {
                val a = structure[i]
                val b = structure[i + 1]
                val direction = a.directionTo(b)
                val distance = a.distanceTo(b)

                for (s in (0..distance)) {
                    when (direction) {
                        Direction.R -> rocks.add(a.copy(x = a.x + s))
                        Direction.L -> rocks.add(a.copy(x = a.x - s))
                        Direction.D -> rocks.add(a.copy(y = a.y + s))
                        Direction.U -> rocks.add(a.copy(y = a.y - s))
                        else -> Unit
                    }
                }

            }
        }
        return rocks
    }

    private fun MutableList<MutableList<Char>>.update(sand: Point) {
        this[sand.y][sand.x] = 'o'
    }

    private fun List<List<Char>>.print() {
        println()
        for (y in this.indices) {
            for (x in this[y].indices) {
                print(this[y][x])
            }
            println()
        }
    }

    private fun List<List<Char>>.count(char: Char): Int {
        var count = 0
        for (y in this.indices) {
            for (x in this[y].indices) {
                if (this[y][x] == char) {
                    count++
                }
            }
        }
        return count
    }

    private fun Char.blocks(): Boolean {
        return this == '#' || this == 'o'
    }

    private fun Point.fallIn(map: List<List<Char>>, minx: Int): Point {
        var rest = false
        var pos = this.copy(x = x - minx)
        while (!rest) {
            if (map[pos.y + 1][pos.x].blocks()) {
                if (map[pos.y + 1][pos.x - 1].blocks()) {
                    if (map[pos.y + 1][pos.x + 1].blocks()) {
                        rest = true
                    } else {
                        pos = Point(pos.x + 1, pos.y + 1)
                    }
                } else {
                    pos = Point(pos.x - 1, pos.y + 1)
                }
            } else {
                pos = Point(pos.x, pos.y + 1)
            }
        }
        return pos
    }
}

fun main() {
    val day = Day14()

    println("Solution1: ${day.solve1()}")
    println("Solution2: ${day.solve2()}")
}