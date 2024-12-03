import { clsx } from "clsx";
import { twMerge } from "tailwind-merge";
import { cubicOut } from "svelte/easing";

/** @param {string[]} inputs */
export function cn(...inputs) {
	return twMerge(clsx(inputs));
}