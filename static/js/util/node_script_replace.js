export function nodeScriptReplace (node) {
  if (nodeScriptIs(node) === true) {
    node.parentNode.replaceChild(nodeScriptClone(node), node)
  } else {
    let i = -1; const children = node.childNodes
    while (++i < children.length) {
      nodeScriptReplace(children[i])
    }
  }

  return node
}
function nodeScriptClone (node) {
  const script = document.createElement('script')
  script.text = node.innerHTML

  let i = -1; const attrs = node.attributes; let attr
  while (++i < attrs.length) {
    script.setAttribute((attr = attrs[i]).name, attr.value)
  }
  return script
}

function nodeScriptIs (node) {
  return node.tagName === 'SCRIPT'
}
