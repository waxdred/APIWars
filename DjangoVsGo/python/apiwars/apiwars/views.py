# myapp/views.py

from prometheus_client import Counter
from uuid import uuid4
from django.http import JsonResponse

# Créer un compteur pour suivre les requêtes
uuid_requests = Counter('uuid_requests_total', 'Total number of UUID requests')

def get_uuid(request):
    uuid_requests.inc()  # Incrémente le compteur pour chaque requête
    return JsonResponse({'uuid': str(uuid4())})
