import json
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt
from django.views.decorators.http import require_http_methods


@csrf_exempt  # Disable CSRF for this demo endpoint
@require_http_methods(["POST"])
def send_data(request):
    try:
        # Parse JSON from request body
        data = json.loads(request.body)
        
        # Print to stdout
        print("Received data:")
        print(json.dumps(data, indent=2))
        
        # Return success response
        return JsonResponse({
            "status": "success",
            "message": "Data received and printed to stdout",
            "received": data
        }, status=200)
        
    except json.JSONDecodeError:
        return JsonResponse({
            "status": "error",
            "message": "Invalid JSON"
        }, status=400)
    except Exception as e:
        return JsonResponse({
            "status": "error",
            "message": str(e)
        }, status=500)