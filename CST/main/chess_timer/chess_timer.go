components {
  id: "chess_timer_gui"
  component: "/main/chess_timer/chess_timer_gui.gui"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "timer_logic"
  component: "/main/chess_timer/timer_logic.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "gui_url"
    value: "#chess_timer_gui"
    type: PROPERTY_TYPE_URL
  }
}
