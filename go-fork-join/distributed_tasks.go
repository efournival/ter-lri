package fj

import (
	"sync"
	"time"
)

type (
	distributedTask struct {
		worker string
		result WorkType
	}

	DistributedTasks struct {
		sync.WaitGroup
		tasks     map[int]*distributedTask
		worker    taskWorker
		workers   []string
		startTime time.Time
		starving  time.Duration
	}
)

/*
	TODO:

	utiliser net/rpc
	utiliser un joli logger
	définir la fonction de travail sur l'objet DistributedTasks
	2 procédures RPC : Work() AckResult()
	lors de fork:
	  - le premier est exécuté en local
	  - à partir du deuxième, logger puis appeller la procédure RPC Work() sur le premier worker et ainsi de suite
	lorsqu'un worker reçoit du travail via Work(), logger puis il effectuer puis appeller AckResult() sur l'émetteur
	AckResult() logger puis assigner le champ result d'une distributedTask et appeller WaitGroup.Done()
	périodiquement appeller une troisième fonction RPC WhatAreYouDoing() renvoyant les identifiants des tâches en cours ?
	si l'appel à WhatAreYouDoing() échoue :
	  - logger
	  - supprimer le worker
	  - renvoyer la tâche
	si la tâche n'est pas dans la réponse de WhatAreYouDoing() et que result est à nil :
	  - mettre un timeOut, le message est sûrement en transit, si il expire :
	  - logger
	  - renvoyer à un autre worker
*/
