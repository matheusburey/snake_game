import pygame
import funcoes
from pygame.locals import *

# Inicia o game
pygame .init()  # inicio o pygame
screen = pygame.display.set_mode((600, 600))  # Tamanho da tela
pygame.display.set_caption('snake')  # Etiqueta do game

# Snake
snake = [(200, 200), (210, 200), (220, 200)]  # Cobra inicial
sn_skin = pygame.Surface((10, 10))  # Tamanho de cada quadrado da cobra
sn_skin.fill((255, 255, 255))  # cor da cobra

# maça
apple_pos = funcoes.on_grid_random()  # Calcula localização aleatoria
apple = pygame.Surface((10, 10))  # Tamanho da maça
apple.fill((255, 0, 0))  # Cor da maça

# Direção
UP = 0
RIGHT = 1
DOWN = 2
LEFT = 3
my_direction = LEFT  # Direção inicial
font = pygame.font.Font('freesansbold.ttf', 18)  # Fonte
clock = pygame.time.Clock()  # Tempo de atualização da tela por segundo
score = 0   # Pontuaçao

while True:
    pygame.display.update()
    clock.tick(10)

    for event in pygame .event.get():
        if event.type == QUIT:
            pygame.quit()
        if event.type == KEYDOWN:
            if event.key == K_UP:
                my_direction = UP
            if event.key == K_DOWN:
                my_direction = DOWN
            if event.key == K_LEFT:
                my_direction = LEFT
            if event.key == K_RIGHT:
                my_direction = RIGHT

    if funcoes.collision(snake[0], apple_pos):
        apple_pos = funcoes.on_grid_random()
        snake.append((0, 0))
        score += 1

    if snake[0][0] == 600 or snake[0][1] == 600 or snake[0][0] < 0 or snake[0][1] < 0:
        break

    if snake[0] in snake[1:]:
        break

    for i in range(len(snake)-1, 0, -1):
        snake[i] = (snake[i-1][0], snake[i-1][1])

    snake.pop()

    if my_direction == UP:
        snake.insert(0,  (snake[0][0], snake[0][1] - 10))
    if my_direction == DOWN:
        snake.insert(0, (snake[0][0], snake[0][1] + 10))
    if my_direction == RIGHT:
        snake.insert(0, (snake[0][0] + 10, snake[0][1]))
    if my_direction == LEFT:
        snake.insert(0, (snake[0][0] - 10, snake[0][1]))

    screen.fill((0, 12, 0))
    screen.blit(apple, apple_pos)

    score_font = font.render('score: %s' % score, True, (155, 255, 255))
    score_rect = score_font.get_rect()
    score_rect.topleft = (600 - 120, 10)
    screen.blit(score_font, score_rect)

    for pos in snake:
        screen.blit(sn_skin, pos)

while True:
    game_over_font = pygame.font.Font('freesansbold.ttf', 75)
    game_over_screen = game_over_font.render('GAME OVER', True, (255, 255, 255))
    game_over_rect = game_over_screen.get_rect()
    game_over_rect.midtop = (600 / 2, 100)
    screen.blit(game_over_screen, game_over_rect)
    pygame.display.update()
    pygame.time.wait(500)
    while True:
        for event in pygame.event.get():
            if event.type == QUIT:
                pygame.quit()
                exit()
