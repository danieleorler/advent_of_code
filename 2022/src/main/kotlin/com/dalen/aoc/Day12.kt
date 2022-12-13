package com.dalen.aoc

import com.dalen.aoc.Utils.getLines

class Day12 {

    fun loadMap(): List<List<Char>> {
        return getLines("12").map { it.toList() }
    }

    fun findLetter(needle: Char, map: List<List<Char>>): Point? {
        for (y in map.indices) {
            for (x in map[y].indices) {
                if (needle == map[y][x]) {
                    return Point(x = x, y = y)
                }
            }
        }
        return null
    }

    fun nextStep(current: Point, map: List<List<Char>>, visited: List<Point>, highestLetterVisited: Char): Point? {
        val currentLetter = if (map.getLetter(current) == 'S') {
            Char('a'.code - 1)
        } else {
            map.getLetter(current)
        }

        val candidateNeighbours = listOf(
            Point(x = current.x - 1, y = current.y),
            Point(x = current.x + 1, y = current.y),
            Point(x = current.x, y = current.y - 1),
            Point(x = current.x, y = current.y + 1),
        )
            .asSequence()
            .filter { it.x >= 0 && it.y >= 0 && it.y < map.size && it.x < map[0].size}
            .filter { !visited.contains(it) }
            .map { PointWithLetter(it, map.getLetter(it)) }

        var bestCandidates = listOf<PointWithLetter>()
        var highest = Char(highestLetterVisited.code)

        while (bestCandidates.isEmpty() && highest >= Char('a'.code - 1)) {
            bestCandidates = candidateNeighbours
                .filter { it.letter.code >= highest.code }
                //.onEach { println("$currentLetter vs $it: ${it.letter} => ${currentLetter.code - it.letter.code}") }
                .filter { currentLetter.code - it.letter.code in -1..0 }
                .toList()

            highest = Char(highest.code - 1)
        }

        if (bestCandidates.isEmpty()) {
            println(highest)
            return candidateNeighbours.filter { it.letter == 'E' }.firstOrNull()?.point
        }

        val candidateLetter = bestCandidates.maxByOrNull { p -> p.letter }!!.letter

        println(candidateLetter)

        return bestCandidates.filter { it.letter == candidateLetter }.sortedByDescending {
            it.point.x - it.point.y
        }.map { it.point }.first()
    }



}

fun main() {
    val day = Day12()

    val map = day.loadMap()

    var current = day.findLetter('S', map)
    val pointsVisited = mutableListOf(current!!)
    var highestLetterVisited = Char('a'.code - 1)

    do {
        println("Current ----> $current on $highestLetterVisited")
        current = day.nextStep(current!!, map, pointsVisited, highestLetterVisited)
        if (current == null) {
            map.printPath(pointsVisited)
        }
        pointsVisited.add(current!!)
        if (map.getLetter(current!!).code > highestLetterVisited.code) {
            highestLetterVisited = map.getLetter(current!!)
        }
    } while (map.getLetter(current!!) != 'E')

    map.printPath(pointsVisited)
    println(pointsVisited.size)
}

fun List<List<Char>>.getLetter(point: Point): Char {
    return this[point.y][point.x]
}

fun List<List<Char>>.printPath(visited: List<Point>) {
    val printable = this.map { it.toMutableList() }

    visited.forEach { printable[it.y][it.x] = '.' }

    for (y in printable.indices) {
        for (x in printable[y].indices) {
            print(printable[y][x])
        }
        println()
    }

}

data class PointWithLetter(
    val point: Point,
    val letter: Char
): Comparable<PointWithLetter> {
    override fun compareTo(other: PointWithLetter): Int {
        return compareValuesBy(this, other, {it.letter}, {it.letter})
    }
}