package com.dalen.aoc

object Utils {
    fun getLines(file: String): List<String> {
        return this::class.java.classLoader.getResourceAsStream(file)
            .bufferedReader()
            .readLines()
    }
}