
import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';

export default class IAccountSettings {

	@prop({ required: true })
	allow_messages_from_unknown: boolean;

	@prop({ required: true })
	email_notifications: boolean;

}
