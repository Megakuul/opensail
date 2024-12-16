/** @param {number} tcc @param {number} boost @returns {string} */
export const getRatingColor = (tcc, boost) => {
  // the tcc anchor is 2.00 because ship tcc's are likely in range 0.00-2.00.
  const tccAnchor = 2.00;
  // move tcc to wrapped range to apply the boost then convert it back to original range (0.00-2.00).
  const boostedTcc = ((tcc - (tccAnchor / 2)) * boost) + (tccAnchor / 2);
  // the color anchor is 150 because hsl 0-150 is red to green.
  const colorAnchor = 150; 
  // move tcc to color range and flip it because red is the lowest and we want lowest to be green.
  const relativeTcc = ((boostedTcc / tccAnchor) * colorAnchor);
  // generate color value by clamping the color range.
  const colorValue = Math.max(0, Math.min(relativeTcc, colorAnchor)); // clamp(anchor - relativeTcc)
  return `hsl(${Math.round(colorValue).toString()} 100% 28%)`;
}