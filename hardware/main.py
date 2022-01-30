# カードタッチの処理のpythonファイル（モーリーが作ってたやつ）
import webbrowser as wb
import pyautogui as pg

# タブを閉じるための関数
def close_tab():
  # Macの場合
  # pg.keyDown('command')
  # pg.press('w')
  # pg.keyUp('command')

  # Windowsの場合
  pg.keyDown('ctrl')
  pg.press('w')
  pg.keyUp('ctrl')

  return 


BASE_URL = "http://localhost:8080/form" # BASE_URLはクライアントのURLって感じ

wb.open(BASE_URL, autoraise=True)


# close_tab()

