import pygame
from random import randint
from pygame.locals import *

status = {
        'snake':{
            'position': [(200, 200), (210, 200), (220, 200)],  # Cobra inicial
            'pixel_size':pygame.Surface((10, 10)),  # Tamanho de cada quadrado da cobra
            'color_skin': ((255, 255, 255))  # cor da cobra: branca
        },
        'apple':{
            'position': [(100, 100)],  # Localização inicial da maçã
            'pixel_size': pygame.Surface((10, 10)),  # Tamanho da maçã
            'color_skin': ((255, 0, 0))   # Cor da maçã
            }
    }

# funçoes
# Retorna Posição aleatoria para maçã 
def on_grid_random():
    x = randint(0, 59)
    y = randint(0, 59)
    return x * 10, y * 10


# Verifiva colisão
def collision(c1, c2):
    return (c1[0] == c2[0]) and (c1[1] == c2[1])


# Desenha a cobra na tela
def print_element(element):
    for po in element['position']:
        pygame.draw.rect(screen, element['color_skin'], [po[0], po[1], 10, 10])


#Verifica tecla clicada
def click_key(key, my_direction):
    for event in key:
        if event.type == QUIT:
            pygame.quit()
        if event.type == KEYDOWN:
            if event.key == K_UP:
                my_direction = 'UP'
            if event.key == K_DOWN:
                my_direction = 'DOWN'
            if event.key == K_LEFT:
                my_direction = 'LEFT'
            if event.key == K_RIGHT:
                my_direction = 'RIGHT'
    return my_direction


def new_position(snake):
    if my_direction == 'UP':
        snake.insert(0,  (snake[0][0], snake[0][1] - 10))
    if my_direction == 'DOWN':
        snake.insert(0, (snake[0][0], snake[0][1] + 10))
    if my_direction == 'RIGHT':
        snake.insert(0, (snake[0][0] + 10, snake[0][1]))
    if my_direction == 'LEFT':
        snake.insert(0, (snake[0][0] - 10, snake[0][1]))


# Inicia o game
pygame.init()  # inicio o pygame
screen = pygame.display.set_mode((600, 600))  # Tamanho da tela
pygame.display.set_caption('snake')  # Etiqueta do game
my_direction = 'LEFT'  # Direção inicial
font = pygame.font.Font('freesansbold.ttf', 18)  # Fonte
clock = pygame.time.Clock()  # Tempo de atualização da tela por segundo
score = 0   # Pontuaçao


while True:
    pygame.display.update()
    clock.tick(10)
    screen.fill((0, 12, 0))
    
    print_element(status['snake'])
    print_element(status['apple'])

    my_direction = click_key(pygame.event.get(), my_direction)

    if collision(status['snake']['position'][0], status['apple']['position'][0]):
        status['apple']['position'][0] = on_grid_random()
        status['snake']['position'].append((0, 0))
        score += 1

    if status['snake']['position'][0][0] == 600 or status['snake']['position'][0][1] == 600 or status['snake']['position'][0][0] < 0 or status['snake']['position'][0][1] < 0:
        break

    if status['snake']['position'][0] in status['snake']['position'][1:]:
        break

    for i in range(len(status['snake']['position'])-1, 0, -1):
        status['snake']['position'][i] = (status['snake']['position'][i-1][0], status['snake']['position'][i-1][1])

    status['snake']['position'].pop()

    new_position(status['snake']['position'])

    score_font = font.render(f'score: {score}', True, (155, 255, 255))
    score_rect = score_font.get_rect()
    score_rect.topleft = (600 - 120, 10)
    screen.blit(score_font, score_rect)


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

