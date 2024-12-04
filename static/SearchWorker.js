// SearchWorker performs a full text search on a large dataset.
// To avoid UI interruption this task is outsourced to a browser worker.

self.onmessage = function(e) {
  try {
    const { components, query } = e.data;

    let searchComponents = JSON.parse(components);
    let matchingComponents = {};
  
    for (const [key, value] of Object.entries(searchComponents)) {
      if (key.includes(query, 0) || JSON.stringify(value).includes(query, 0)) {
        matchingComponents[key] = value;
      }
    }
    self.postMessage({status: "", matching: matchingComponents});
  } catch (err) {
    self.postMessage({status: err.message, matching: {}});
  }
}