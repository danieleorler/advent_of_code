package com.dalen.aoc

import com.dalen.aoc.Utils.getLines
import java.lang.Integer.max
import java.lang.Integer.min
import kotlin.math.abs
import kotlin.math.sign

class Day09 {

    private val instructionRegex = """^([RLUD])\s(\d+)$""".toRegex()


    fun solve1(): Int {
        val rope = Rope(
            head = Point(0, 0),
            knots = mutableMapOf(
                1 to Knot(Point(0, 0), true)
            )
        )

        return simulate(rope)
    }

    fun solve2(): Int {
        val rope = Rope(
            Point(0, 0),
            mutableMapOf(
                1 to Knot(Point(0, 0)),
                2 to Knot(Point(0, 0)),
                3 to Knot(Point(0, 0)),
                4 to Knot(Point(0, 0)),
                5 to Knot(Point(0, 0)),
                6 to Knot(Point(0, 0)),
                7 to Knot(Point(0, 0)),
                8 to Knot(Point(0, 0)),
                9 to Knot(Point(0, 0), true),
            )
        )

        return simulate(rope)
    }

    private fun simulate(rope: Rope): Int {
        val visited = mutableSetOf(rope.head)
        getLines("09")
            .map { instructionRegex.find(it) }
            .map {
                val (d, s) = it!!.destructured
                val direction = Direction.valueOf(d)
                val steps = s.toInt()

                executeInstruction(direction, steps, rope)
            }
            .forEach { visited.addAll(it) }

        return visited.size
    }

    private fun executeInstruction(direction: Direction, steps: Int, rope: Rope): MutableSet<Point> {
        val visited = mutableSetOf<Point>()
        for (step in (0 until steps)) {
            rope.head = rope.head.move(direction)

            var head = rope.head
            var knotId = 1

            do {
                var knot = rope.knots[knotId]!!
                if (!areTouching(head, knot.point)) {
                    knot = knot.follow(head)
                    if (knot.isTail) {
                        visited.add(knot.point)
                    }
                }
                rope.knots[knotId] = knot
                head = knot.point
                knotId += 1
            } while (rope.knots.containsKey(knotId))
        }

        return visited
    }

    private fun areTouching(head: Point, tail: Point): Boolean {
        return max(abs(head.x - tail.x), abs(head.y - tail.y)) < 2
    }

    private fun Knot.follow(head: Point): Knot {
        val xDiff = point.x - head.x
        val yDiff = point.y - head.y

        val xDirection = if (xDiff.sign != 0) xDiff.sign * -1 else 0
        val yDirection = if (yDiff.sign != 0) yDiff.sign * -1 else 0

        return copy(point = Point(point.x + min(abs(xDiff), 1) * xDirection, point.y + min(abs(yDiff), 1) * yDirection))
    }
}

data class Rope(
    var head: Point,
    val knots: MutableMap<Int, Knot>
)

data class Knot(
    val point: Point,
    val isTail: Boolean = false
)

fun main() {
    val day = Day09()

    println("Solution1: ${day.solve1()}")
    println("Solution2: ${day.solve2()}")
}
