package handlers

import (
	"context"
	"log"
	contexts "wrench/app/contexts"
	settings "wrench/app/manifest/action_settings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SnsPublishHandler struct {
	Next           Handler
	ActionSettings *settings.ActionSettings
}

type SnsActions struct {
	SnsClient *sns.Client
}

func (handler *SnsPublishHandler) Do(ctx context.Context, wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext) {
	if wrenchContext.HasError == false {
		settings := handler.ActionSettings.SNS
		message := bodyContext.GetBodyString()
		actor := new(SnsActions)

		publishInput := sns.PublishInput{TopicArn: aws.String(settings.TopicArn), Message: aws.String(message)}

		groupId := getGroupId(settings.GroupId, wrenchContext, bodyContext)
		if groupId != "" {
			publishInput.MessageGroupId = aws.String(groupId)
		}

		// if dedupId != "" {
		// 	publishInput.MessageDeduplicationId = aws.String(dedupId)
		// }
		// if filterKey != "" && filterValue != "" {
		// 	publishInput.MessageAttributes = map[string]types.MessageAttributeValue{
		// 		filterKey: {DataType: aws.String("String"), StringValue: aws.String(filterValue)},
		// 	}
		// }
		_, err := actor.SnsClient.Publish(ctx, &publishInput)
		if err != nil {
			log.Printf("Couldn't publish message to topic %v. Here's why: %v", settings.TopicArn, err)
		}
	}

	if handler.Next != nil {
		handler.Next.Do(ctx, wrenchContext, bodyContext)
	}
}

func (handler *SnsPublishHandler) SetNext(next Handler) {
	handler.Next = next
}

func getGroupId(groupIdValue string, wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext) string {
	var groupId string
	if len(groupIdValue) > 0 {
		if contexts.IsCalculatedValue(groupIdValue) {
			command := contexts.ReplaceCalculatedValue(groupIdValue)
			if contexts.IsWrenchContextCommand(command) {
				groupId = contexts.GetValueWrenchContext(command, wrenchContext)
			} else {
				groupId = contexts.GetValueBodyContext(command, bodyContext)
			}
		} else {
			groupId = groupIdValue
		}
	}

	return groupId
}
