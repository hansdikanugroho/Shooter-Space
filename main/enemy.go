components {
  id: "enemy"
  component: "/main/enemy.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"enemy\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 279.2\n"
  "  y: 266.6\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/sprites.atlas\"\n"
  "}\n"
  ""
  rotation {
    z: 0.9999889
    w: 0.004708291
  }
  scale {
    x: 0.225925
    y: 0.239718
  }
}
embedded_components {
  id: "hit area"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 1.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemy\"\n"
  "mask: \"bullet\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"box\"\n"
  "  }\n"
  "  data: 31.698233\n"
  "  data: 32.86046\n"
  "  data: 10.0\n"
  "}\n"
  "bullet: true\n"
  ""
}
